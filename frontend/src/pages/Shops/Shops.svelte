<script lang="typescript">
    import {get} from 'svelte/store';
    import {_} from 'svelte-i18n'

    import {LoadingStatus} from '../../stores/utils';
    import {
        shops,
        shopsTotal,
        shopsStatus,
        allShops,
        updateCurrentShopsSlice
    } from '../../stores/shops';
    import {EntityType, ActionType} from '../../stores/utils';

    import {style} from "./style";

    import {mockData} from "./data";
    import {Shop} from "./types";

    import CommonTable
        from '../../common/components/Tables/CommonTable/CommonTable.svelte'
    import CommonMap
        from '../../common/components/Map/CommonMap/CommonMap.svelte'
    import {CommonTable as CommonTableInterface} from '../../common/components/Tables/CommonTable/utils'
    import SimpleTextElement
        from '../../common/components/ElementsAndBlocks/SimpleTextElement/SimpleTextElement.svelte'
    import MapActionElement
        from '../../common/components/ActionElements/MapActionElement/MapActionElement.svelte'
    import ShopActionsElement
        from './ShopActionsElement/ShopActionsElement.svelte'
    import {ShopActionsData} from './ShopActionsElement/utils';
    import {maxRecordsPerPage} from '../../common/components/Pagination/PaginationBlock/utils';
    import ButtonIconEdit
        from '../../common/components/Buttons/ButtonIconEdit/ButtonIconEdit.svelte'
    import ButtonIconRemove
        from '../../common/components/Buttons/ButtonIconRemove/ButtonIconRemove.svelte'
    import MapContainer
        from '../../common/components/Map/MapContainer/MapContainer.svelte'
    import UrlElement
        from '../../common/components/ElementsAndBlocks/UrlElement/UrlElement.svelte'
    import {onMount} from 'svelte';
    import EditShopActionElement
        from "./EditShopActionElement/EditShopActionElement.svelte";

    const onPageChange = async (currentPage: number) => {
        // updateCurrentShopsSlice((currentPage - 1) * maxRecordsPerPage, currentPage * maxRecordsPerPage - 1)
        await updateCurrentShopsSlice((currentPage - 1) * maxRecordsPerPage, maxRecordsPerPage)
    }

    const tableData: CommonTableInterface = {
        title: "shops.titles.shops",
        total: 0,
        withButton: true,
        buttonTitle: "shops.buttons.new",
        buttonClickHandler: () => {
        },
        status: LoadingStatus.Loading,
        columnsConfig: [
            {
                header: 'common.labels.name',
                component: UrlElement,
                overflowed: true,
                style: 'flex: 1 0 20%',
                mapping: (data: Shop) => {
                    return {
                        name: data.name,
                        url: data.url,
                    }
                }
            },
            {
                header: 'common.labels.address',
                component: SimpleTextElement,
                style: 'flex: 1 0 20%',
                mapping: (data: Shop) => data.address,
            },
            {
                header: 'common.labels.comment',
                component: SimpleTextElement,
                style: 'flex: 1 0 40%',
                mapping: (data: Shop) => data.comment,
            },
            {
                header: "",
                component: ShopActionsElement,
                style: 'flex: 1 0 10%',
                mapping: (data: Shop): ShopActionsData => {
                    return {
                        shopData: data,
                        mapData: [{
                            name: data.name,
                            address: data.address,
                            id: data.id,
                            actionType: ActionType.Map,
                            entityType: EntityType.Shop,
                        }],
                    }
                },
            },
        ],
        data: []
    }

    onMount(async () => {
        shopsTotal.set(mockData.length);

        await onPageChange(1);
        tableData.status = get(shopsStatus);
        tableData.total = get(shopsTotal);
        tableData.data = get(shops);
    });
</script>

<div>
    <CommonTable onPageChange={onPageChange} withButton={tableData.withButton}
                 buttonTitle={tableData.buttonTitle} status={tableData.status}
                 total={tableData.total} data={$shops}
                 config={tableData.columnsConfig} title={tableData.title}>
        <div slot="titleButton">
            <EditShopActionElement withButton={true}
                                   buttonTitle={$_("shops.buttons.new")}/>
        </div>
    </CommonTable>
    <CommonMap status={tableData.status}
               data={$shops.map(elem => ({name: elem.name, address: elem.address}))}/>
</div>