import React from "react";
import "./style.css"

interface info{
    infoName:string,
    tips:string
}

export default function (infos:info){
    return (
        <div className={"quest"}>
            <span>{infos.infoName}: </span><input className={"infoType"} placeholder={infos.tips}></input>.
        </div>
    )
}