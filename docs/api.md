# MiRolls API

MiRolls has some apis, if you want to develop a MiRolls theme, you must use them. They have some data specification, and something you can set in **mt.json**.

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
    options: string[/*put your options in there*/];
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
It is ```questionnaire```, we have 6 types of questions. But some types' underlying code look same like ```input``` and ```textarea```
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

Every questionnaire has two value, they are ```link``` and ```code```, but they won't in ```questionnaire``` structure, they will return as response in [make questionnaire api](#Make Questionnaire)
