import config from 'config';
import { authHeader } from '../_helpers';

export const userService = {
    login,
    logout,
    register,
    getAll,
    getPlans,
    getPubKey
};

function login(username, password) {

    const requestOptions = {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({user : {
            email: username,
            password: password,
        }})
    };
    
    return fetch(`${config.apiUrl}/api/users/login`, requestOptions)
        .then(handleResponse)
        .then(user => {
  
            // login successful if there's a jwt token in the response
            if (user.user.token) {
                // store user details and jwt token in local storage to keep user logged in between page refreshes
                localStorage.setItem('user', JSON.stringify(user.user));
            }
            return user;
        });
}


function logout() {
    // remove user from local storage to log user out
    localStorage.removeItem('user');
}

function register(user) {
    const requestOptions = {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({user:user})
    };
    console.log(requestOptions.body)
    return fetch(`${config.apiUrl}/api/users`, requestOptions)
        .then(handleResponse)
        .then(user => {
        if (user.user.token) {
            localStorage.setItem('user', JSON.stringify(user.user));
        }
        return user;
    });;
}

function getAll() {
    const requestOptions = {
        method: 'GET',
        headers: authHeader()
    };
    return fetch(`${config.apiUrl}/api/movie/all`, requestOptions).then(handleResponse);
}

function getPlans() {
    const requestOptions = {
        method: 'GET',
    };
    return fetch(`${config.apiUrl}/api/stripe/plan`, requestOptions).then(handleResponse);
}

function getPubKey() {
    const requestOptions = {
        method: 'GET',
    };
    return fetch(`${config.apiUrl}/api/stripe/pubkey`, requestOptions).then(handleResponse);
}


function handleResponse(response) {
    return response.text().then(text => {
        const data = text && JSON.parse(text);
        if (!response.ok) {
            const error = (data && data.message) || response.statusText;
            return Promise.reject(error);
        }
        console.log(data)
        return data;
    });
}