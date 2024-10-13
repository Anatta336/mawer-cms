# Family History CMS
Will be a specialised content management system for working with articles about history of families and places.

## Status
Inoperative.

## Plan

### Parse `.docx` to create article
- Basic form done, although no footnotes or images yet.

### Article content
- Articles made up of a very minimal set of possible items.
- Represent each paragraph and heading as an independent thing - will use that when generating vector embeddings.

### Images
- Positioned in article using float left, float right, or block.
- Multiple images in a single block?
- Automatic re-sizing and re-encoding. Make sure we have original version (or a lossless conversion) available.
- Click image to view full size.

### Footnotes
- Expand out either entirely inline, or after the current paragraph.
    - https://fivethirtyeight.com/features/what-does-an-0-7-start-tell-you-about-an-nfl-coach/

### People
- Records in database, which can be linked back to articles.
- Will have a generated page for each person, with links to articles that mention them.
- Could have LLM summary generated. Suspect better to have that hand-written, although could be a starting point.

### Search
- Try that contextual retrieval stuff for generating good embeddings for each paragraph, group or paragraphs, or article as a whole.
    - https://www.anthropic.com/news/contextual-retrieval
- Rather than true RAG, just want to use the retrieval part to surface information.
- Although also try sticking it in the context of an LLM and see what the result is.

## Requirements
- What form is the "content" that'll be shown? I'd guess a mixture of images, scanned documents (maybe PDF files?), and your own writing. Is your writing text with basic formatting, or do you weave in tables and other more complex structures?
    - Word, pdf, jpeg, PPt files. All my written text starts off in Word.
- Are there any unusual files you find you're dealing with, that could do with being put on the site somehow?
    - Things like birth certificates are always in pdf or jpeg form.
- Is the plan to have the content managed by you, or should other (trusted) people be able to log in and add/edit things too?
    - Just me!
- Do you have ideas in terms of organisation? Perhaps a family tree as a starting point? I'm imagining a page for each "item" (a person, a place, an organisation, etc.) that will each have links to related items.
    - Front page with images for people to click onto. One House History. One Driver and Weatherhead. Then there will be other family ones eg Campbell/ Bedford etc. Behind each would sit a page with introductory text, documents to download, a gallery, hyperlinks.
- Are there any examples of other sites or other media that'd act as inspiration for the look and feel you'd like to have on the site?
    - Not really, but I use Braunston Village website. Although I have not updated  it this year I edit the Friendly Club pages. Braunston Friendly Club - Braunston, Northamptonshire (https://www.braunston.org.uk/org.aspx?n=Braunston-Friendly-Club) I have just added a note on my to do list to update this!!!
- As a goal is long term archiving of data, I'm thinking it'll be valuable to provide a one-stop "download everything" button so visitors can grab all the raw data and keep a store of it themselves. We can set up automated backups too, but flexibility and easy access to the complete dataset is probably the best hope for this kind of potentially multi-generational data storage.
