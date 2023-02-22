import "./global.css"
import Steps from "../Steps";
import React, {FormEvent, useState} from "react";
import "../ServerInfo/style.css"
import axios from "axios";
import {Navigate, useLocation} from "react-router-dom";
import InputBar from "../InputBar";


export default function (){
    interface siteInfo {

        name: string,
        link: string,
        logo: string,
        icp: string,
        lang: string,
        needIcp: number,
        mainColor:string

        [key: string]: string | number;
    }



    const [siteInfo,setSiteInfo] = useState<siteInfo>({
        name: "",
        link: "",
        logo: "",
        icp: "A Nice Questionnaire System",
        lang: "en",
        needIcp: 0,
        mainColor:"rgb(21, 127, 248)"
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
            alert("Error! Please restart the program." + err)
        })
    }

    return (
        <div className={"page"} style={{opacity:isOpacity}}>
            <Steps step={3} CompletedStep={0} nowStep={1}></Steps>
            <h1>Enter Your Server Info</h1>
            <InputBar placeholder={"MiRolls"} question={"Your Site Name: "}></InputBar>
            <InputBar placeholder={'"zh" or "en"'} question={"Your Size Language: "}></InputBar>
            <InputBar placeholder={"https://github.com/fluidicon.png"} question={"Your Size Logo: "}></InputBar>
            <span className={"footerTips"}>There all content correspond config.yaml</span>
            <button className={"nextStep"} onClick={submit}>{buttonText}</button>
            {isGoNext}
        </div>
    )
}