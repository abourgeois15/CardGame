import React from 'react';
import {BrowserRouter, Route, Routes} from 'react-router-dom';
import MenuPage from './page/Menu.page';
import GamePage from './page/Game.page';


export default function App() {
  return (
      <BrowserRouter>
          <Routes>
            <Route path="/" element={<MenuPage/>}/>
            <Route path="/game" element={<GamePage/>}/>
          </Routes>
      </BrowserRouter>
  );
}
