import "./style.css";
import React, {useEffect, useState} from 'react';

interface items{
    message:string,
    type:"error"|"success"|"warning",
    title:string
    onCloseToDo?:Function
}

function Index(props:items) {
    const [changeCount,setChangeCount] = useState(0);
    const [message,setMessage] = useState(props.message);

    function defaultFunc():boolean{
        return false
    }
    const onClose:Function = props.onCloseToDo != undefined ?
        props.onCloseToDo as Function :
        defaultFunc

    const vmRender = <div className={"dialogWindow"}>
        <div className={"dialog"}>
            <span className={"dialogTitle " + props.type}>
                <i className="fa-solid fa-circle-exclamation"></i>
                <span className={"title"}>{props.title}</span>
            </span>
            <span className={"dialogControl"} onClick={()=>{setChangeCount(0);setMessage("");onClose()}}>
                    <i className="fa-sharp fa-solid fa-xmark"></i>
                </span>
            <div className={"content"}>
                {message}
            </div>
        </div>
    </div>


    useEffect(()=>{
        setChangeCount(changeCount + 1)
    },[message]);

    useEffect(()=>{
        setMessage(props.message)
    },[props.message])

    function renderJsx():JSX.Element{
        if(changeCount == 2){
            return vmRender
        }else {
            return <></>
        }
    }

    return renderJsx()
}

export default Index;