import Steps from "../Steps";
import {useState} from "react";
import axios from "axios";
import {Navigate} from "react-router-dom";

export default function (){
    const [downloadSpeed,setDownloadSpeed] = useState("")
    const [done,setDone] = useState(<></>)
    const [opacity,setOpacity] = useState(1)

    function download(){
        setDownloadSpeed("Downloading ~")
        axios.post("/install/download",{file:"default"}).then(res=>{
            if(res.data.message === "success"){
                setDownloadSpeed("Download Done!")
                setTimeout(()=>{
                    setOpacity(0)
                    setTimeout(()=>{
                        setDone(<Navigate to={"/done"}></Navigate>)
                    },1000)
                },1000)
            }else {
                setDownloadSpeed("Can't download file." + res.data.error)
            }
        }).catch(err => {
            setDownloadSpeed("Can't download file." + err)
        })
    }

    return (
        <div className={"page"} style={{opacity}}>
            <Steps step={3} CompletedStep={2} nowStep={3}></Steps>
            <h1>Choice a theme</h1>
            <p className={"tips"}>MiRolls before 2.0.0 can only use the default theme. terribly sorry!</p>
            <button className={"nextStep"} onClick={download}>Download default theme</button>
            <p className={"tips"}>{downloadSpeed}</p>
            {done}
        </div>
    )
}