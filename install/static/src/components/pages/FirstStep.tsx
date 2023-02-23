import "./global.css"
import Steps from "../Steps";
import React, {useRef, useState} from "react";
import "../ServerInfo/style.css"
import axios from "axios";
import {Navigate} from "react-router-dom";
import InputBar from "../InputBar";
import mode from "../../mode";


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

    const [buttonText,setButtonText] = useState("Click To Submit")
    const [isOpacity,setIsOpacity] = useState(1)
    const [isGoNext,setIsGoNext] = useState(<></>)
    const siteName = useRef(null);
    const siteLang = useRef(null);
    const siteLogo = useRef(null);

    function submit(){
        const devMode = mode.debug

        if (devMode === mode.debug){
            console.log(getData())
        }else{
            setIsOpacity(0)
            setButtonText("Submitting Data...")
            axios.post("/install/set/site",getData()).then(res=>{
                setIsGoNext(<Navigate to={"/step-two"}/>)
                return res.data
            }).catch(err=>{
                alert("Error! Please restart the program." + err)
            })
        }
    }

    function getData():siteInfo{
    // function getData():{}{
        // return {
        //     name:
        // }
        return {
            name: (siteName.current as any).getAnswer(),
            link: getRootUrl(),
            logo: (siteLogo.current as any).getAnswer(),
            icp: "A Nice Questionnaire System",
            lang: (siteLang.current as any).getAnswer(),
            needIcp: 0,
            mainColor:"rgb(21, 127, 248)"
        }
    }

    function getRootUrl():string {
        const re = new RegExp(/^.*\//);
        return (re.exec(window.location.href) as string[])[0];
    }

    return (
        <div className={"page"} style={{opacity:isOpacity}}>
            <Steps step={3} CompletedStep={0} nowStep={1}></Steps>
            <h1>Enter Your Server Info</h1>
            <InputBar placeholder={"MiRolls"} question={"Your Site Name: "} ref={siteName}></InputBar>
            <InputBar placeholder={'"zh" or "en"'} question={"Your Size Language: "} ref={siteLang}></InputBar>
            <InputBar placeholder={"https://github.com/fluidicon.png"} question={"Your Size Logo: "} ref={siteLogo}></InputBar>
            <span className={"footerTips"}>There all content correspond config.yaml</span>
            <button className={"nextStep"} onClick={submit}>{buttonText}</button>
            {isGoNext}
        </div>
    )
}