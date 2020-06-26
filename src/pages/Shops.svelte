<script lang="typescript">
    import { LoadingStatus } from '../stores/utils';
    
    import { mockData } from "./data";
    import { Shop } from "./types";

    import CommonTable from '../common/components/CommonTable/CommonTable.svelte'
    import { CommonTable as CommonTableInterface } from '../common/components/CommonTable/utils'
    import SimpleTextElement from '../common/components/SimpleTextElement/SimpleTextElement.svelte'
    import MapActionElement from '../common/components/MapActionElement/MapActionElement.svelte'
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
          header: "Actions",
          component: MapActionElement,
          style: 'flex: 1 0 10%',
          mapping: (data: Shop) => [{
              name: data.name, 
              address: data.address
            }],
        }
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
</div>