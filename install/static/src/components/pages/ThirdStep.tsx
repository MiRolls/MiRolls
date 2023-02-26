import Steps from "../Steps";

export default function (){
    function download(){

    }

    return (
        <div className={"page"}>
            <Steps step={3} CompletedStep={2} nowStep={3}></Steps>
            <h1>Choice a theme</h1>
            <p className={"tips"}>MiRolls before 2.0.0 can only use the default theme. terribly sorry!</p>
            <button className={"nextStep"} onClick={download}>Download default theme</button>
        </div>
    )
}