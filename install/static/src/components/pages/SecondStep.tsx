import Steps from "../Steps";
import React, {useRef} from "react";
import InputBar from "../InputBar";

export default function (){
    const dbUser = useRef(null)
    const dbPwd = useRef(null)
    const dbServer = useRef(null)
    const dbName = useRef(null)

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
        </div>
    )
}