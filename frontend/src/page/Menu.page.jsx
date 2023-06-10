import React, { useState } from "react";
import Stack from '@mui/material/Stack';
import Button from '@mui/material/Button';
import BasePage from "../component/BasePage";
import Slider from "@mui/material/Slider";
import Typography from "@mui/material/Typography";
import { FormControl } from "@mui/material";
import { httpStatus } from "../config";
import { createGame } from "./gameApi";
import { useNavigate } from "react-router-dom";

const MenuPage = () => {
    const navigate = useNavigate()
    const [nbPlayers, setNbPlayers] = useState(4)

    const handleNewGame = () => {
        createGame(nbPlayers, (error, result) => {
            if (!error && result && result.status === httpStatus.OK) {
                console.log(result)
            }
        })
        navigate("/game")
    }
    return (
        <BasePage >
            <Stack 
                spacing={2}
                direction="column"
                style={{alignItems: "center"}}
            >
                <FormControl>
                    <Typography>
                        Choose The Number of Players
                    </Typography>
                    <Slider
                        value={nbPlayers}
                        valueLabelDisplay="auto"
                        onChange={(event, newValue) => {
                            setNbPlayers(newValue)
                        }}  
                        step={1}
                        min={2}
                        max={8}
                        style={{
                            width: 500
                        }}
                    />
                </FormControl>
                
                <Button 
                    variant="contained"
                    style={{
                        height: 100,
                        width: 500,
                        fontSize: 40
                    }}
                    onClick={handleNewGame}
                >
                    Start New Game
                </Button>
            </Stack>
        </BasePage>
    )
}

export default MenuPage