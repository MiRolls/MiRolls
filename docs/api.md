# MiRolls Docs

MiRolls has some apis, if you want to develop a MiRolls theme, you must use them. They have some data specification, and
something you can set in **mt.json**.

## Data

MiRolls(normal) has three structures, they are ```bigData``` ```questionnaire``` and ```answer```

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

### Create Questionnaire

Request:

A [```questionnaire```](#Questionnaire)

Response:

| Key  | Value                | Description                                                                            |
|------|----------------------|----------------------------------------------------------------------------------------|
| code | 16 random characters | use md5 generate, must show user it, you can get bigdata/answers/questionnaire with it |
| link | 16 random characters | use md5 generate, must show user it, you can answer questionnaire with it              |

What do the api will do?
1. Check whether the data is ```questionnaire```
2. Generate code and link
3. Write database and return