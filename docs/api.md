# MiRolls API

MiRolls has some apis, if you want to develop a MiRolls theme, you must use them. They have some data specification, and something you can set in **mt.json**.

## Data

MiRolls(normal) has two structures, they are ```questionnaire``` and ```answer```

```typescript
// questionnaire

interface Questionnaire {
    title: string;
    // the title of questionnaire
    questions: (RadioCheckboxAndSelect | InputAndtextarea | Slider)[]
}

interface Question {
    title: string;
    type: "radio" | "checkbox" | "input" | "textarea" | "select" | "slider";
    // use tag name, more easy for develop
}

interface RadioCheckboxAndSelect extends Question {
    type: "radio" | "checkbox" | "select";
    options: string[/*put your options in there*/];
}

interface InputAndtextarea extends Question {
    type: "input" | "textarea",
    placeholder: string
    // <input placeholder="" type="text">
}

interface Slider extends Question {
    type: "slider",
    range: [number, number],
    unit: number,
}
```
