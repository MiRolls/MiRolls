# MiRolls Docs

MiRolls has some apis, if you want to develop a MiRolls theme, you must use them. They have some data specification.

## Data

All APIs revolve around four structures

They are ```bigData``` ```site``` ```questionnaire``` and ```answer```

### Site

If you want to develop a MiRolls theme, you have to create a file, mt.json

If you use Vite with React/Vue, it looks like that

```
your-project
├─ src
│   └─ your-code
├─ public
│    └─ mt.json
├─ dist
|   ├─ index.html
|   └─ mt.json
```

**mt.json must in your theme root path.**

A qualified mt.json file looks like this

``````json
{
  "name": "sllor",
  "not-found-page": "index.html",
  "frame": "react",
  "config": {
    "site": {
      "name": {
        "type": "string",
        "value": "string",
        "placeholder": "Sllor",
        "tag": "input",
        "option-name": "Site Name"
      },
      "site-url": {
        "type": "string",
        "value": "url",
        "placeholder": "https://www.example.com/",
        "tag": "input",
        "tips": "Your target domain / IP",
        "option-name": "Domain/IP"
      },
      "main-color": {
        "type": "string",
        "value": "select",
        "placeholder": [
          "tomato",
          "red",
          "ruby",
          "crimson",
          "pink",
          "plum",
          "purple",
          "violet",
          "iris",
          "indigo",
          "blue",
          "cyan",
          "teal",
          "jade",
          "green",
          "grass",
          "brown",
          "orange",
          "sky",
          "mint",
          "lime",
          "yellow",
          "amber",
          "gold",
          "bronze",
          "gray"
        ],
        "option-name": "Main Color"
      },
      "logo": {
        "type": "string",
        "value": "image",
        "option-name": "Site Logo",
        "editor": "update"
      },
      "introduce": {
        "type": "string",
        "value": "markdown",
        "placeholder": "# Sllor /n A good and pretty site",
        "option-name": "Introduce",
        "tips": "[MARKDOWN] Your text will show in index page"
      },
      "about": {
        "type": "string",
        "value": "markdown",
        "placeholder": "# About \n > An apple ......",
        "tag": "textarea",
        "editor": "markdown",
        "option-name": "About",
        "tips": "[MARKDOWN] Your text will show in a dictionary, /about"
      },
      "footer": {
        "type": "string",
        "value": "markdown",
        "placeholder": "## Sllor \n [About](/about) | [......",
        "tag": "textarea",
        "edit": "markdown",
        "option-name": "Footer",
        "tips": "[MARKDOWN] Your text will show at the end of every page."
      }
    }
  }
}
``````

The ```site``` structure is in ```config``` key, you can DIY it. You can get your data via [get site api](#get-site)

### Questionnaire

```typescript
// questionnaire

interface Questionnaire {  //main interface
    title: string;
    // the title of questionnaire
    questions: (RadioCheckboxAndSelect | InputAndtextarea | Slider)[];
}

interface Question {
    title: string;
    type: "radio" | "checkbox" | "input" | "textarea" | "select" | "slider";
    // use tag name, more easy for develop
}

interface RadioCheckboxAndSelect extends Question {
    type: "radio" | "checkbox" | "select";
    options: string[];
    // options look like["orange","apple","banana"]
}

interface InputAndtextarea extends Question {
    type: "input" | "textarea";
    placeholder: string;
    // <input placeholder="" type="text">
}

interface Slider extends Question {
    type: "slider";
    range: [number, number];
    unit: number;
}
```

It is ```questionnaire```, we have 6 types of questions. But some types' underlying code look same like ```input```
and ```textarea```
```questionnaire```is a very important in MiRolls program.

If you want to develop a MiRolls Theme, you should pay attention different question types.

### Answer

```typescript
interface Answer {
    link: string;
    answers: string | boolean[]
    /* 
    * If a user create a multiple choice
    * - [x] Yes
    * - [ ] No
    * The answers should looks like this:
    * [true,false]
    */
}

```

Every questionnaire has two value, they are ```link``` and ```code```, but they won't in ```questionnaire``` structure,
they will return as response in [create questionnaire api](#create-questionnaire)

## Apis

**All MiRolls apis can only be request via POST**

### Normal data

| Key     | Type                  | Description                             |
|---------|-----------------------|-----------------------------------------|
| message | "success" \| "error   | api running status                      |
| error   | string                | if message="error", there is error info |
| data    | any, usually {} or [] | response data                           |

### Site

#### Get Site

Request:

```/site/get```

Response (data):

[```Site```](#site)

What do the api will do?

1. Return ```site``` to you

### Questionnaire

#### Create Questionnaire

Request:

A [```questionnaire ```](#Questionnaire)structure.

Response (data):

| Key  | Value                | Description                                                                            |
|------|----------------------|----------------------------------------------------------------------------------------|
| code | 16 random characters | use md5 generate, must show user it, you can get bigdata/answers/questionnaire with it |
| link | 16 random characters | use md5 generate, must show user it, you can answer questionnaire with it              |

What do the api will do?
1. Check whether the data is ```questionnaire```
2. Generate code and link
3. Write database and return


