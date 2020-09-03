import { SvelteComponentDev } from 'svelte/internal';
import { LoadingStatus } from '../../../../stores/utils';

export interface CommonTable {
    title: string;
    columnsConfig: ColumnConfig[];
    data: Record<any, any>[]
    total: number;
    status: LoadingStatus;
    withButton: boolean;
    buttonTitle: string;
    overflowed?: boolean;
    buttonClickHandler: () => void;
}

export interface ColumnConfig {
[x: string]: any;
    header: string;
    component: SvelteComponentDev;
    style?: string;
    mapping: (data: Record<any, any>) => any;
}
