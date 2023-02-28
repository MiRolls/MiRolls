import Steps from "../Steps";
import React, {useRef, useState} from "react";
import InputBar from "../InputBar";
import axios from "axios";
import Dialog from "../Dialog";
import {Navigate} from "react-router-dom";

export default function () {
    interface dbConfig {
        Username: string,
        Password: string,
        Protocol: string,
        Host: string,
        Port: number,
        Database: string,
    }

    const [isGoNext, setIsGoNext] = useState(<></>)
    const [isOpacity, setIsOpacity] = useState(1)
    const [dialogErrorMessage,setDialogErrorMessage] = useState("")
    const dbUser = useRef(null);
    const dbPwd = useRef(null);
    const dbServer = useRef(null);
    const dbName = useRef(null);

    function submit() {
        setIsOpacity(0)
        axios.post("/install/set/database", getDbConfig()).then(res => {
            if (res.data.message === "success") {
                setIsGoNext(<Navigate to="/step-three"></Navigate>)
            } else {
                setDialogErrorMessage("Server has error! Please reboot Mirolls-program. " + res.data.error)
            }
        }).catch(err => {
            setDialogErrorMessage("Server has error! Please reboot Mirolls-program. " + err)
        })
    }

    function getDbConfig(): dbConfig {
        return {
            Username: (dbUser.current as any).getAnswer(),
            Password: (dbPwd.current as any).getAnswer(),
            Protocol: "tcp",
            Host: (dbServer.current as any).getAnswer(),
            Port: 3306,
            Database: (dbServer.current as any).getAnswer()
        }
    }

    function goHome(){
        setIsGoNext(<Navigate to={"/"}></Navigate>)
    }

    return (
        <div>
            <Dialog message={dialogErrorMessage} type={"error"} title={"Error!"} onCloseToDo={goHome}></Dialog>
            <div className={"page"} style={{opacity: isOpacity}}>
                <Steps step={3} CompletedStep={1} nowStep={2}></Steps>
                <h1>Setting Your Database</h1>
                <InputBar placeholder={"MiRolls"} question={"Your Database Username: "} ref={dbUser}/>
                <InputBar placeholder={"password112233"} question={"Your Database Password: "} ref={dbPwd}></InputBar>
                <InputBar placeholder={"localhost"} question={"Your Database Server: "} ref={dbServer}></InputBar>
                <InputBar placeholder={"mirolls"} question={"Your Database Name: "} ref={dbName}></InputBar>
                <span className={"tips"}>This step is very important. If this info is false, MiRolls will have many error.</span>
                <span className={"tips"}>If you need other settings, you can modify config.yaml</span>
                <button className={"nextStep"} onClick={submit}>Click To Submit</button>
                {isGoNext}
            </div>
        </div>
    )
}