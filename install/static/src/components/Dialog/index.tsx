import "./style.css";
import React, {useEffect, useState} from 'react';

interface items{
    message:string,
    type:"error"|"success"|"warning",
    title:string
}

function Index(props:items) {

    const vmRender = <div className={"dialogWindow"}>
        <div className={"dialog"}>
                <span className={"dialogTitle " + props.type}><i className="fa-solid fa-circle-exclamation"></i>
                    <span className={"title"}>{props.title}</span>
                </span>
            <div className={"content"}>
                {props.message}
            </div>
        </div>
    </div>
    const [changeCount,setChangeCount] = useState(0)

    useEffect(()=>{
        setChangeCount(changeCount + 1)
    },[props.message]);

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