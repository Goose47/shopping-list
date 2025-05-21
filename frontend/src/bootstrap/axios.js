import axios from 'axios';

const setupAxios = () => {
    window.axios = axios;
    window.axios.defaults.headers.common['Telegram-Init'] = window.Telegram.WebApp.initData;
    axios.defaults.baseURL = 'https://api.rwfshr.ru';
}

export {setupAxios}