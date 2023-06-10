import { API_URL, DefaultHeaders, HTTPGet, HTTPPost } from "../config";

// Request employee by its id. Call end-point GET /employees/{EmployeeID}
const createGame = async (nb, cb) => {
    try {
        const response = await fetch(`${API_URL}/game/${nb}`, {
            method: HTTPPost,
            headers: DefaultHeaders,
        });
        return cb(null, {body: await response.json(), status: response.status});
    } catch (error) {
        return cb(error, null);
    }
};

// Request employee by its id. Call end-point GET /employees/{EmployeeID}
const getGame = async (cb) => {
    try {
        const response = await fetch(`${API_URL}/game`, {
            method: HTTPGet,
            headers: DefaultHeaders,
        });
        return cb(null, {body: await response.json(), status: response.status});
    } catch (error) {
        return cb(error, null);
    }
};

export {
    createGame,
    getGame
}