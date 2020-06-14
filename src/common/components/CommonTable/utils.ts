import { SvelteComponentDev } from 'svelte/internal';

export interface CommonTable {
    title: string;
    columnsConfig: ColumnConfig[];
    data: Record<any, any>[]
    total: number;
}

export interface ColumnConfig {
    header: string;
    component: SvelteComponentDev;
    style?: string;
}
