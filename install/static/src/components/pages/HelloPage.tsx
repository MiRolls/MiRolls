import {Navigate} from "react-router-dom";
import "./guide.css"
import "./global.css"
import {useState} from "react";
import axios from "axios";
import Dialog from "../Dialog";

export default function (){
    const [isRenderNext,setRenderNext] = useState(false)
    const [isOpacity,setIsOpacity] = useState(1)
    const [dialogMessage,setDialogMessage] = useState("This is a dialog.")

    function goStepOne(){
        if(isRenderNext){
            return <Navigate to="/step-one"/>
        }else{
            return
        }
    }

    function renderAnimation(){
        setIsOpacity(0)
        setTimeout(()=>{
            setRenderNext(true)
        },500)
    }

    return(
        <div className={"page"} style={{opacity:isOpacity}}>
            <img src={"/favicon.png"} className={"logo"} alt={"MiRolls Install Program"}/>
            <h1 id={"title"}>MiRolls <span className={"installText"}>Install</span> Program</h1>
            <button onClick={()=>renderAnimation()} className={"startSetup"}>Start Setup</button>
            {goStepOne()}
            <button className={"startSetup"} onClick={()=>setDialogMessage("There is change")}>Render!</button>
            <Dialog message={dialogMessage} type={"error"} title={"Error!"}></Dialog>
        </div>
    )
}