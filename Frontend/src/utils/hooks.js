import { MobXProviderContext } from 'mobx-react';
import { useContext } from 'react';
import Store from '../store';

export function useStores(name) {
    const stores = useContext(MobXProviderContext);
    return stores[name];
}

