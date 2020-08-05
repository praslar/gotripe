export function authHeader() {
    // return authorization header with jwt token
    let user = JSON.parse(localStorage.getItem('user'));
    console.log("debug here", user)
    if (user.token) {
        return { 'Authorization': user.token };
    } else {
        return {};
    }
}