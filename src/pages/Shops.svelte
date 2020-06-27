<script lang="typescript">
    import { LoadingStatus } from '../stores/utils';
    
    import { SideMainPadding } from "./style";

    import { mockData } from "./data";
    import { Shop } from "./types";

    import CommonTable from '../common/components/CommonTable/CommonTable.svelte'
    import CommonMap from '../common/components/CommonMap/CommonMap.svelte'
    import { CommonTable as CommonTableInterface } from '../common/components/CommonTable/utils'
    import SimpleTextElement from '../common/components/SimpleTextElement/SimpleTextElement.svelte'
    import MapActionElement from '../common/components/MapActionElement/MapActionElement.svelte'
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
          component: MapActionElement,
          style: 'flex: 1 0 10%',
          mapping: (data: Shop) => [{
              name: data.name, 
              address: data.address
            }],
        },
        // {
        //   header: "",
        //   component: ButtonIconEdit,
        //   style: 'flex: 1 0 10%',
        //   mapping: (data: Shop) => [{
        //       name: data.name, 
        //       address: data.address
        //     }],
        // },
        // {
        //   header: "",
        //   component: ButtonIconRemove,
        //   style: 'flex: 1 0 10%',
        //   mapping: (data: Shop) => [{
        //       name: data.name, 
        //       address: data.address
        //     }],
        // }
      ],
      data: []
    }

    onMount(async () => {
        setTimeout(() => {
          tableData.status = LoadingStatus.Finished;
          // tableData.data = [];
          // tableData.total = 0;
          tableData.total = 30;
          tableData.data = mockData;
        }, 5000)
    });
</script>

<style lang="scss">
</style>

<div>
    <CommonTable withButton={tableData.withButton} buttonTitle={tableData.buttonTitle} status={tableData.status} total={tableData.total} data={tableData.data} config={tableData.columnsConfig} title={tableData.title} />
    <CommonMap status={tableData.status} data={tableData.data.map(elem => ({name: elem.name, address: elem.address}))}/>
</div>