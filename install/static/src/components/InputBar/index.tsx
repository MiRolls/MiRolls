import React, {useImperativeHandle} from 'react';
import "./style.css"

interface input{
    placeholder:string,
    question:string,
}

function Index(props:input) {
    // useImperativeHandle(()=>{
    //
    // })
    return (
        <div className={"quest"}>
            <span>{props.question}</span>
            <input className={"infoType"} placeholder={props.placeholder}></input>.
        </div>
    );
}

export default Index;