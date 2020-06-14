<script lang="typescript">
    import { LoadingStatus } from '../stores/utils';

    import CommonTable from '../common/components/CommonTable/CommonTable.svelte'
    import { CommonTable as CommonTableInterface } from '../common/components/CommonTable/utils'
    import SimpleElement from '../common/components/SimpleElement/SimpleElement.svelte'
    import { onMount } from 'svelte';

    const tableData: CommonTableInterface = {
      title: "Shops",
      total: 0,
      status: LoadingStatus.Loading,
      columnsConfig: [
        {
          header: 'id',
          component: SimpleElement,
          style: 'flex: 1 0 30%'
        },
        {
          header: 'name',
          component: SimpleElement,
          style: 'flex: 1 0 30%'
        },
        {
          header: 'number',
          component: SimpleElement,
          style: 'flex: 1 0 30%'
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
          tableData.data = [
            { id: 1, name: 'name1', number: 1 },
            { id: 2, name: 'name2', number: 2 },
            { id: 3, name: 'name3', number: 3 },
            { id: 4, name: 'name4', number: 4 },
            { id: 5, name: 'name5', number: 5 },
          ]
        }, 5000)
    });
</script>

<style lang="scss">
</style>

<div>
    <CommonTable status={tableData.status} total={tableData.total} data={tableData.data} config={tableData.columnsConfig} title={tableData.title} />
</div>