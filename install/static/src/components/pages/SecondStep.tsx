import Steps from "../Steps";
import React, {useRef, useState} from "react";
import InputBar from "../InputBar";
import axios from "axios";

export default function (){
    interface dbConfig{
        username:string,
        password:string,
        protocol:string,
        host:string,
        port:number,
        database:string,
    }

    const dbUser = useRef(null);
    const dbPwd = useRef(null);
    const dbServer = useRef(null);
    const dbName = useRef(null);
    // refs

    const [isGoNext,setIsGoNext] = useState(<></>)

    function submit(){
        axios.post("/install/set/database",getDbConfig()).then(res => {

        })
    }

    function getDbConfig():dbConfig{
        return {
            username:(dbUser.current as any).getAnswer(),
            password:(dbPwd.current as any).getAnswer(),
            protocol:"tcp",
            host:(dbServer.current as any).getAnswer(),
            port:3306,
            database:(dbServer.current as any).getAnswer()
        }
    }

    return (
        <div className={"page"}>
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
    )
}