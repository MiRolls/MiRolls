import Steps from "../Steps";
import {useState} from "react";
import axios from "axios";
import {Navigate} from "react-router-dom";

export default function (){
    const [downloadSpeed,setDownloadSpeed] = useState("")
    const [done,setDone] = useState(<></>)
    const [opacity,setOpacity] = useState(1)

    let getDownloadControl: number | undefined;
    function getDownloadSpeed(){
        setOpacity(0)
        axios.post("/install/download/speed").then(res=>{
            if(res.data.message != "success"){
                setDownloadSpeed("Can't get download speed." + res.data.error);
            }else {
                setDownloadSpeed(res.data.data.speed as string + "mb/s");
                if (res.data.down){
                    clearInterval(getDownloadControl);

                }
            }
        }).catch(err => {
            setDownloadSpeed("Can't get download speed.." + err)
            setDone(<Navigate to={"/done"}></Navigate>)
        })
    }

    function download(){
        axios.post("/install/download",{file:"default"}).then(res=>{
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