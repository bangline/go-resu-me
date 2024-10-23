# Go-Resu-Me

Super simple résumé builder using Go, Tailwind & good ole YAML.

**Features**
- Based on the excellent template from [Gergely Orosz on the Pragmatic Engineer site](https://blog.pragmaticengineer.com/the-pragmatic-engineers-resume-template/). Unfortunately, I don't use Google Docs or Word and I personally found fiddling in Apple Pages just an awful experience. I was learning Go so I built this tool.
- Clean print styles for exporting to PDF (just save it from your browser)
- Markdown syntax for entries (very limited though)
- Colors can easily be customised
- Simple custom themes

## Usage

```bash
Usage go-resu-me:
  -a string
    	Color of the accent text (default "#2e1065")
  -bg string
    	Color of the web background (default "#f5f3ff")
  -h string
    	Color of the heading text (default "#0a0a0a")
  -out string
    	Path to output html file for the resume (default "resume.html")
  -source string
    	Path to yaml file to build resume from (default "resume.yml")
  -t string
    	Color of the non heading text (default "#0c0a09")
  -theme string
    	Theme to use for the resume (default "default")
```

Generate the HTML and view it in a browser. [Tailwindcss](https://tailwindcss.com/) is used liberally in the templates and included in the page from the source `cdn.tailwindcss.com`.

If you just want to run go-resu-me locally you can use

## YAML Structure

Everything in the yaml is a string, dates, and numbers are parsed as strings. Bullet points support a very limited subset of markdown syntax

- `**Bold**`
- `[Hyperlinks](https://bangline.dev)`

```yaml
---
name: David Kennedy
details:
  - Glasgow, UK
  - david@bangline.co.uk
  - +44(0)123456789
  - "[bangline.co.uk](http://bangline.dev)"
personal_statement: >
  Hello my name is Dave and this can spread over...
  Multiple lines
sections:
  - title: Work experiences
    entries:
      # Simple work experience with only 1 position held at a company
      - type: work
        position: Robot 1
        company: Acme
        dates: Summer - Winter
        bullets:
          - bullet 1
          - bullet 2
      # A more complex work experience with multiple positions held at the same company
      - type: work
        company: Blahmo
        dates: May - December
        roles:
          - position: Robot 2
            dates: November - December
            team: A-Team
            bullets:
              - bullet 3
              - bullet 4
              - bullet 5
          - position: Robot 3
            dates: June - November
            team: B-Team
            bullets:
              - bullet 6
  - title: Education
    entries:
      - **My Degree** My university
      - School
  - title: Personal interests
    entries:
      - Sleeping
      - Eating
      - Coding
      - In that order
```

## Themes

The template files are super easy to customise. So if you want different markup, in the [templates](./templates/) directory simply make a copy of the `default` folder and edit away.

> [!WARNING]
> The flag `-theme` passed to the CLI is simply used to specify the directory the templates are loaded from. You have to keep the same file names.
