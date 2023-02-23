import React, {forwardRef, useImperativeHandle, useState} from 'react';
import "./style.css"

interface input{
    placeholder:string,
    question:string,
}

export default forwardRef((props:input,ref) => {
    const [answer,setAnswer] = useState("")
    useImperativeHandle(ref,()=>({
        getAnswer: ():string=>{
            return answer
        }
    }))
    return (
        <div className={"quest"}>
            <span>{props.question}</span>
            <input className={"infoType"} placeholder={props.placeholder} onInput={(event)=>setAnswer(event.currentTarget.value)}></input>.
        </div>
    );
});

