import "./global.css"
import Steps from "../Steps";
import React, {useEffect} from "react";

export default function (){
    return (
        <div className={"page"}>
            <Steps step={3} CompletedStep={0}></Steps>
            <h1>Server</h1>
        </div>
    )
}