import React from 'react';
import "./global.css"
import Steps from "../Steps";

function Done() {
    return (
        <>
            <div className={"page"}>
                <Steps CompletedStep={3} step={3}></Steps>
                <img src={"/favicon.png"} className={"logo"} alt={"MiRolls Install Program"}/>
                <h1>You completed the install of MiRolls!</h1>
            </div>
        </>
    );
}

export default Done;