# shelfpace

> **I'm learning Go by building this.** It's my first real Go project. I build it in 2-hour blocks, and every block ships something I can link to. The [ledger](LEDGER.md) shows each block: what shipped, which Go concepts it taught, and what I got wrong.

shelfpace turns a reading backlog into a daily plan you can actually keep. You give it your books, the themes you care about (with weights), and your book-club meeting dates. It gives back a day-by-day plan sized to the minutes you really have, delivered to your phone as a subscribed calendar.

## Why I'm building this

I love reading and I'm bad with reading management. I am part of a book club I love, I love meeting someone who looks at me reading and suggests something and I love books of people I admire. All of these scenarios have really blown up my backlog and I need to be better. I also want to learn Go so this made sense.

## How it will work

- **Deadlines first.** Book-club meetings are hard deadlines, scheduled earliest-deadline-first. If a meeting can't be met at your real pace, you get told how far short you'll fall and what would fix it.
- **Themes get fair shares.** Minutes are split across themes in proportion to their weights (stride scheduling), with a cap on how many books are open at once.
- **It learns your pace.** Each logged session updates a running average of pages per minute and of how much of the plan you actually read.
- **A missed day is not debt.** Tomorrow is re-planned from your current page. Nothing stacks up.
- **No AI.** The "intelligence" is a deterministic, testable scheduling algorithm.

## Status

Not started. Block 1 (`shelfpace pace`, a pages-per-day calculator for one book and one deadline) is next. See [LEDGER.md](LEDGER.md).

## Non-goals

Single user, self-hosted. No social features, ratings or reviews (that's what Goodreads, StoryGraph and Hardcover are for). No accounts, and no one else's credentials.

## License

[MIT](LICENSE)
