import Steps from "../Steps";
import {useState} from "react";
import axios from "axios";

export default function (){
    const [downloadSpeed,setDownloadSpeed] = useState("")

    let getDownloadControl: number | undefined;
    function getDownloadSpeed(){
        axios.post("/install/download/speed").then(res=>{
            if(res.data.message != "success"){
                setDownloadSpeed("Can't get download speed." + res.data.error);
            }else {
                setDownloadSpeed(res.data.data.speed as string + "mb/s");
                if (res.data.down === "yes"){
                    clearInterval(getDownloadControl);
                }
            }
        }).catch(err => {
            setDownloadSpeed("Can't get download speed.." + err)
        })
    }

    function download(){
        axios.post("/install/download/default").then(res=>{
            if(res.data.message === "success"){
                getDownloadControl = setInterval(()=>{
                    getDownloadSpeed()
                })
            }else {
                setDownloadSpeed("Can't download file." + res.data.error)
            }
        }).catch(err => {
            setDownloadSpeed("Can't download file." + err)
        })
    }

    return (
        <div className={"page"}>
            <Steps step={3} CompletedStep={2} nowStep={3}></Steps>
            <h1>Choice a theme</h1>
            <p className={"tips"}>MiRolls before 2.0.0 can only use the default theme. terribly sorry!</p>
            <button className={"nextStep"} onClick={download}>Download default theme</button>
            <p className={"tips"}>{downloadSpeed}</p>
        </div>
    )
}