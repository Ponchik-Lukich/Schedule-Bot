# Schedule Bot
Bot for finding free offices of the MEPhI Research Institute.
Bot is available on: [@mephi_schedule_bot](https://t.me/mephi_checker_bot)

Bot is serverless function, which is deployed on Yandex Cloud Functions.

## Local Setup

To run the bot locally, you will first need to set up your environment:

1. Clone this repository:

```bash
git clone <repository_url>
cd path_to_repository
```

2. Install the required dependencies:

```bash
go mod download
```

3. Setup environment variables:

- `TELEGRAM_BOT_TOKEN`: Your Telegram bot token.
- `DSN`: Your PostgreSQL connection string.

4. Run the bot:

```bash
go run cmd/main.go
```