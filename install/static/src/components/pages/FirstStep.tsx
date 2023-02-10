import "./global.css"
import Steps from "../Steps";
import React, {FormEvent, useState} from "react";
import "../ServerInfo/style.css"


export default function (){
    interface siteInfo {

        name: string;
        link: string;
        logo: string;
        mainColor: string;
        icp: string;
        lang: string;
        needIcp: boolean;
        [key: string]: any;
    }


    const [siteInfo,setSiteInfo] = useState<siteInfo>({
        name: "",
        link: "",
        logo: "",
        mainColor:"",
        icp:"",
        lang:"",
        needIcp: false,
    })

    function changeInfo(options:string){
        let siteInfo2 = siteInfo
        return function (val:FormEvent<HTMLInputElement>){
            siteInfo2.site[options] = val.currentTarget.value
            setSiteInfo(siteInfo2)
        }
    }

    return (
        <div className={"page"}>
            <Steps step={3} CompletedStep={0}></Steps>
            <h1>Enter Your Server Info</h1>
            <div className={"quest"}>
                <span>Your Site Name: </span>
                <input className={"infoType"} placeholder="MiRolls" onInput={changeInfo("name")}></input>.
            </div>
        </div>
    )
}