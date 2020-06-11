import { SvelteComponent } from 'svelte'

export interface CommonTable {
    title: string;
    columns: ColumnConfig[];
    data: Record<any, any>[]
}

export interface ColumnConfig {
    header: string;
    data: any;
    component: SvelteComponent;
}
