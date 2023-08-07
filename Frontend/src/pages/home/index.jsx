import React from 'react'
import { observer } from 'mobx-react'
import { useStores } from 'utils/hooks'

const Home = observer((props) => {
    console.log('Home', props);
    const token = useStores('token');
    console.log('store', token.getToken());
    return (
        <div>Home</div>
    )
})
export default Home
