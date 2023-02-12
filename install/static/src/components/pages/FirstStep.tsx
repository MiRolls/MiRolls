import "./global.css"
import Steps from "../Steps";
import React, {FormEvent, useState} from "react";
import "../ServerInfo/style.css"
import axios from "axios";
import {Navigate} from "react-router-dom";


export default function (){
    interface siteInfo {

        name: string;
        link: string;
        logo: string;
        icp: string;
        lang: string;
        needIcp: boolean;
        [key: string]: string | boolean;
    }



    const [siteInfo,setSiteInfo] = useState<siteInfo>({
        name: "",
        link: "",
        logo: "",
        icp: "",
        lang: "",
        needIcp: false,
    })
    const [buttonText,setButtonText] = useState("Click To Submit")
    const [isOpacity,setIsOpacity] = useState(1)
    const [isGoNext,setIsGoNext] = useState(<></>)

    function changeInfo(options:string){
        let siteInfo2 = siteInfo
        return function (val:FormEvent<HTMLInputElement>){
            siteInfo2[options] = val.currentTarget.value
            setSiteInfo(siteInfo2)
        }
    }

    function submit(){
        setIsOpacity(0)
        setButtonText("Submitting Data...")
        axios.post("///",siteInfo).then(res=>{
            setIsGoNext(<Navigate to={"/step-two"}/>)
            return res.data
        }).catch(err=>{
            alert("Error! Please restart the program. " + err)
        })
    }

    return (
        <div className={"page"} style={{opacity:isOpacity}}>
            <Steps step={3} CompletedStep={0}></Steps>
            <h1>Enter Your Server Info</h1>
            <div className={"quest"}>
                <span>Your Site Name: </span>
                <input className={"infoType"} placeholder="MiRolls" onInput={changeInfo("name")}></input>.
            </div>
            <div className={"quest"}>
                <span>Your Site Language: </span>
                <input className={"infoType"} placeholder='"zh" or "en"' onInput={changeInfo("name")}></input>.
            </div>
            <div className={"quest"}>
                <span>Your Site Language: </span>
                <input className={"infoType"} placeholder='"zh" or "en"' onInput={changeInfo("name")}></input>.
            </div>
            <span className={"footerTips"}>There all content correspond config.yaml</span>
            <button className={"nextStep"} onClick={submit}>{buttonText}</button>
            {/*<span className={"footerTips"}>The details database config can change config.yaml file</span>*/}
            {isGoNext}
        </div>
    )
}