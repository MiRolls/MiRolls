import "./global.css"
import Steps from "../Steps";
import React, {useImperativeHandle} from "react";
import ServerInFo from "../ServerInfo";

export default function (_: any, ref: React.Ref<unknown> | undefined){
    useImperativeHandle(ref,()=>({

    }))

    return (
        <div className={"page"}>
            <Steps step={3} CompletedStep={0}></Steps>
            <h1>Enter Your Server Info</h1>
            <ServerInFo infoName={"Your Site Name"} tips={"Such of Leason's WebSite"}></ServerInFo>
        </div>
    )
}