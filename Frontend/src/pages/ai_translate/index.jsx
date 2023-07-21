import React, {useState} from 'react'
import axios from 'axios'
import {Input, Button} from 'antd'


export default function Index() {
    const [text, setText] = useState('')
    const [translateText, setTranslateText] = useState('')

    function isJSON(str) {
        if (typeof str == 'string') {
            try {
                JSON.parse(str);
                return true;
            } catch(e) {
                console.log(e);
                return false;
            }
        }  
    }

    const translate = (text) => {
        axios.post(
            '/ai/api/iplayground/rest/v1/flow/454/build/predict',
            {
                message: text
            },
            {
                ContentType: 'application/json'
            }
        ).then(res => {
            console.log(res.data.result)
            let str = (res.data.result).replace(/(```)|(json)/g, '')
            if (isJSON(str)) {
                let obj = JSON.parse(str)
                setTranslateText(obj.result || '')
            } else {
                setTranslateText(str.result || '')
            }
        }).catch(err=> {
            console.log(err)
        })
    }


    return (
        <div>
            <h1>AI翻译?</h1>
            <Input onChange={(e) => setText(e.target.value)} />
            <Button onClick={() => translate(text)}>翻译</Button>
            <Input value={translateText} />
        </div>
    )
}
