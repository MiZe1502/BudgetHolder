import { SvelteComponentDev } from 'svelte/internal';
import { LoadingStatus } from '../../../stores/utils';

export interface CommonTable {
    title: string;
    columnsConfig: ColumnConfig[];
    data: Record<any, any>[]
    total: number;
    status: LoadingStatus;
}

export interface ColumnConfig {
    header: string;
    component: SvelteComponentDev;
    style?: string;
}
