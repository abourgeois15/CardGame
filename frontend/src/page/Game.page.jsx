import React, { useEffect, useState } from "react"
import BasePage from "../component/BasePage"
import { getGame } from "./gameApi"
import { httpStatus } from "../config"


const GamePage = () => {
    const [game, setGame] = useState()

    const loadGame = () => {
        getGame((error, result) => {
            if (!error && result && result.status === httpStatus.OK) {
                console.log(result)
                setGame(result.body.payload)
            }
        })
    }

    useEffect(loadGame, [])
    return (
        <BasePage>

        </BasePage>
    )
}

export default GamePage;