// 将参数转换成功 formdata 接收格式
function stringify(data = {}) {
    const formData = new FormData();
    for (const key in data) {
        // eslint-disable-next-line no-prototype-builtins
        if (data.hasOwnProperty(key)) {
            if (data[key]) {
                if (data[key].constructor === Array) {
                    if (data[key][0]) {
                        if (data[key][0].constructor === Object) {
                            formData.append(key, JSON.stringify(data[key]));
                        } else {
                            data[key].forEach((item, index) => {
                                formData.append(key + `[${index}]`, item);
                            });
                        }
                    } else {
                        formData.append(key + '[]', '');
                    }
                } else if (data[key].constructor === Object) {
                    formData.append(key, JSON.stringify(data[key]));
                } else {
                    formData.append(key, data[key]);
                }
            } else if (data[key] === 0) {
                formData.append(key, 0);
            } else {
                formData.append(key, '');
            }
        }
    }
    return formData;
}
// 防抖
function debounce(fn, delay) {
    let timer = null;
    return function () {
        if (timer) {
            clearTimeout(timer);
        }
        timer = setTimeout(fn, delay);
    };
}

// sessionStorage保存token
function setToken(token) {
    sessionStorage.setItem('USER_TOKEN', token);
}
function getToken() {
    return sessionStorage.getItem('USER_TOKEN');
}
function removeToken() {
    sessionStorage.removeItem('USER_TOKEN');
}

export {
    stringify,
    debounce,
    setToken,
    getToken,
    removeToken
}