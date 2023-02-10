import "./step.css"
import Step from "../Step";

interface props{
    step:number
    CompletedStep:number
}

export default function (props:props){
    const steps = Array(props.step).fill("BigClever")

    const stepList = steps.map((_,index:number)=> (
        <Step hasBeenCompleted={((index + 1) <= props.CompletedStep)} stepNumber={index+1} key={index+233}></Step>
    ))
    return (
        <div className="stepIndicator">
            {stepList}
            {/*<Step hasBeenCompleted={true} stepNumber={1}></Step>*/}
            {/*<Step hasBeenCompleted={false} stepNumber={2}></Step>*/}
            {/*<Step hasBeenCompleted={false} stepNumber={3}></Step>*/}
        </div>
    )
}