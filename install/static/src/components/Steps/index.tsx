import "./step.css"
import Step from "../Step";

interface props{
    step:number
    CompletedStep:number
    nowStep:number
}

export default function (props:props){
    const steps = Array(props.step).fill("MiRollsAuthor:liangmi")

    const stepList = steps.map((_,index:number)=> {
        let i
        if(index+1 === props.nowStep){
            i = 2
        }else if( index+1 <= props.CompletedStep ){
            i = 1
        }else{
            i = 0
        }
        // console.log(i)
        return (
            // <Step hasBeenCompleted={(index + 1) <= props.nowStep} stepNumber={index+1} key={index+233}></Step>
            <Step hasBeenCompleted={i} stepNumber={index+1} key={index+233}></Step>
        )
    })
    return (
        <div className="stepIndicator">
            {stepList}
            {/*<Step hasBeenCompleted={true} stepNumber={1}></Step>*/}
            {/*<Step hasBeenCompleted={false} stepNumber={2}></Step>*/}
            {/*<Step hasBeenCompleted={false} stepNumber={3}></Step>*/}
        </div>
    )
}