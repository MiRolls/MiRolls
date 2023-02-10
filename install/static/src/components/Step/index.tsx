import "./style.css"

interface props{
    hasBeenCompleted:boolean
    stepNumber:number,
}

export default function (steps:props){
    return (
       <div className={"stepCircle " + ((steps.hasBeenCompleted) ? "active" : "unActive")}>{ steps.stepNumber }</div>
    )
}