import "./style.css"
import {useState} from "react";

interface props{
    hasBeenCompleted:number
    stepNumber:number,

}

export default function (steps:props){
    const [circleColor,setCircleColor] = useState("active")

    function getCircleColor(){
        if(!steps.hasBeenCompleted){
            //not completed
            setCircleColor("unActive")
        }else if (steps.hasBeenCompleted === 1){
            setCircleColor("active")
        }else{
            setCircleColor("willActive")
        }
        return circleColor
    }

    return (
       <div className={"stepCircle " + circleColor}>{ steps.stepNumber }</div>
    )
}