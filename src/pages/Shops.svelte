<script lang="typescript">
    import { get } from 'svelte/store';

    import { LoadingStatus } from '../stores/utils';
    import { shops, shopsTotal, shopsStatus } from '../stores/shops';
    
    import { SideMainPadding } from "./style";

    import { mockData } from "./data";
    import { Shop } from "./types";

    import CommonTable from '../common/components/CommonTable/CommonTable.svelte'
    import CommonMap from '../common/components/CommonMap/CommonMap.svelte'
    import { CommonTable as CommonTableInterface } from '../common/components/CommonTable/utils'
    import SimpleTextElement from '../common/components/SimpleTextElement/SimpleTextElement.svelte'
    import MapActionElement from '../common/components/MapActionElement/MapActionElement.svelte'
    import ShopActionsElement from '../common/components/ShopActionsElement/ShopActionsElement.svelte'
    import { ShopActionsData } from '../common/components/ShopActionsElement/utils';
    import ButtonIconEdit from '../common/components/Buttons/ButtonIconEdit/ButtonIconEdit.svelte'
    import ButtonIconRemove from '../common/components/Buttons/ButtonIconRemove/ButtonIconRemove.svelte'
    import MapContainer from '../common/components/MapContainer/MapContainer.svelte'
    import UrlElement from '../common/components/UrlElement/UrlElement.svelte'
    import { onMount } from 'svelte';

    const tableData: CommonTableInterface = {
      title: "Shops",
      total: 0,
      withButton: true,
      buttonTitle: "Add new shop",
      buttonClickHandler: () => {},
      status: LoadingStatus.Loading,
      columnsConfig: [
        {
          header: 'Name',
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
          header: 'Address',
          component: SimpleTextElement,
          style: 'flex: 1 0 20%',
          mapping: (data: Shop) => data.address,
        },
        {
          header: 'Comment',
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
                address: data.address
              }],
            }
          },
        },
      ],
      data: []
    }

    onMount(async () => {
        setTimeout(() => {
          shopsStatus.set(LoadingStatus.Finished);
          shops.set(mockData);
          shopsTotal.set(30);


          tableData.status = get(shopsStatus);
          // tableData.data = [];
          // tableData.total = 0;
          tableData.total = get(shopsTotal);
          tableData.data = get(shops);
        }, 5000)
    });
</script>

<style lang="scss">
</style>

<div>
    <CommonTable withButton={tableData.withButton} buttonTitle={tableData.buttonTitle} status={tableData.status} total={tableData.total} data={$shops} config={tableData.columnsConfig} title={tableData.title} />
    <CommonMap status={tableData.status} data={$shops.map(elem => ({name: elem.name, address: elem.address}))}/>
</div>