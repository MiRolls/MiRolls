import "./style.css"
import {useState} from "react";

interface props{
    hasBeenCompleted:number
    stepNumber:number,

}

export default function (steps:props){
    let circleColor = "active"

    function getCircleColor(){
        if(!steps.hasBeenCompleted){
            //not completed
            circleColor = "unActive"
        }else if (steps.hasBeenCompleted === 1){
            circleColor = "active"
        }else{
            circleColor = "willActive"
        }
        return circleColor
    }

    return (
       <div className={"stepCircle " + getCircleColor()}>{ steps.stepNumber }</div>
    )
}