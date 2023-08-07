import {makeAutoObservable} from 'mobx';


class Token {
    token = 'test'
    constructor() {
        makeAutoObservable(this);
        this.token;
    }
    
    setToken(token) {
        this.token = token;
    }
    getToken() {
        return this.token;
    }

}

export default new Token();