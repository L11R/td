package telegram_test

import (
	"context"
	"crypto/rand"
	"errors"
	"os"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"go.uber.org/zap/zaptest"
	"golang.org/x/sync/errgroup"

	"github.com/gotd/td/telegram"
	"github.com/gotd/td/telegram/internal/e2etest"
	"github.com/gotd/td/telegram/internal/tgtest"
	"github.com/gotd/td/tg"
	"github.com/gotd/td/transport"
)

func testTransportAttempt(ctx context.Context, t *testing.T, trp telegram.Transport) error {
	t.Helper()

	log := zaptest.NewLogger(t)
	defer func() { _ = log.Sync() }()

	client := telegram.NewClient(telegram.TestAppID, telegram.TestAppHash, telegram.Options{
		Addr:      telegram.AddrTest,
		Transport: trp,
	})

	if err := client.Connect(ctx); err != nil {
		return err
	}

	defer func() {
		_ = client.Close()
	}()

	if err := telegram.NewAuth(
		telegram.TestAuth(rand.Reader, 2),
		telegram.SendCodeOptions{},
	).Run(ctx, client); err != nil {
		return err
	}

	if _, err := client.Self(ctx); err != nil {
		return err
	}

	return nil
}

func testTransport(trp telegram.Transport) func(t *testing.T) {
	return func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
		defer cancel()

		var err error

		// Sometimes testing server can return "AUTH_KEY_UNREGISTERED" error.
		// It is expected and client implementation is unlikely to cause
		// such errors, so just doing retries with sleep.
		for i := 0; i < 10; i++ {
			if err = testTransportAttempt(ctx, t, trp); err == nil {
				return // ok
			}

			var rpcErr *telegram.Error
			if errors.As(err, &rpcErr) {
				switch rpcErr.Type {
				case "NEED_MEMBER_INVALID", "AUTH_KEY_UNREGISTERED":
					time.Sleep(time.Second)
					continue
				}
			}

			break // unrecoverable error
		}

		t.Fatal(err)
	}
}

func TestExternalE2EConnect(t *testing.T) {
	if ok, _ := strconv.ParseBool(os.Getenv("GOTD_TEST_EXTERNAL")); !ok {
		t.Skip("Skipped. Set GOTD_TEST_EXTERNAL=1 to enable external e2e test.")
	}

	t.Run("Abridged", testTransport(transport.Abridged(nil)))
	t.Run("Intermediate", testTransport(transport.Intermediate(nil)))
	t.Run("PaddedIntermediate", testTransport(transport.PaddedIntermediate(nil)))
	t.Run("Full", testTransport(transport.Full(nil)))
}

const dialog = `— Да?
— Алё!
— Да да?
— Ну как там с деньгами?
— А?
— Как с деньгами-то там?
— Чё с деньгами?
— Чё?
— Куда ты звонишь?
— Тебе звоню.
— Кому?
— Ну тебе.`

func TestE2EUsersDialog(t *testing.T) {
	if ok, _ := strconv.ParseBool(os.Getenv("GOTD_USERS_DIALOG")); !ok {
		t.Skip("Skipped. Set GOTD_USERS_DIALOG=1 to enable users dialog e2e test.")
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()
	log := zaptest.NewLogger(t).WithOptions(zap.IncreaseLevel(zapcore.InfoLevel))

	cfg := e2etest.TestConfig{
		AppID:   telegram.TestAppID,
		AppHash: telegram.TestAppHash,
		DcID:    2,
		Addr:    telegram.AddrTest,
	}
	suite := e2etest.NewSuite(tgtest.NewSuite(ctx, t, log), cfg, rand.Reader)

	g, gctx := errgroup.WithContext(ctx)
	auth := make(chan *tg.User, 1)
	botCtx, botCancel := context.WithCancel(gctx)
	g.Go(func() error {
		return e2etest.NewEchoBot(suite, auth).Run(botCtx)
	})

	user := <-auth
	g.Go(func() error {
		defer botCancel()
		return e2etest.NewUser(suite, strings.Split(dialog, "\n"), user).Run(gctx)
	})

	require.NoError(t, g.Wait())
}
