# Choose your own adventure



## Exercise details

Choose Your Own Adventure is a series of books where as you read you would occasionally be given options about how you want to proceed. The goal of this exercise is to recreate this experience via a web application where each page will be a portion of the story, and at the end of every page the user will be given a series of options to choose from (or be told that they have reached the end of that particular story arc).

Stories will be provided via a JSON file with the following format:

```json
{
  "story-arc": {
    "title": "A title for that story arc. Think of it like a chapter title.",
    "story": [
      "A series of paragraphs, each represented as a string in a slice.",
      "This is a new paragraph in this particular story arc."
    ],
    // Options will be empty if it is the end of that
    // particular story arc. Otherwise it will have one or
    // more JSON objects that represent an "option" that the
    // reader has at the end of a story arc.
    "options": [
      {
        "text": "the text to render for this option. eg 'venture down the dark passage'",
        "arc": "the name of the story arc to navigate to. This will match the story-arc key at the very root of the JSON document"
      }
    ]
  },
  ...
}
```

