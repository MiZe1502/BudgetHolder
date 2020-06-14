<script lang="typescript">
    import CommonTable from '../common/components/CommonTable/CommonTable.svelte'
    import { CommonTable as CommonTableInterface } from '../common/components/CommonTable/utils'
    import SimpleElement from '../common/components/SimpleElement/SimpleElement.svelte'
	  import { onMount } from 'svelte';

    const tableData: CommonTableInterface = {
      title: "Shops",
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
    <CommonTable data={tableData.data} config={tableData.columnsConfig} title={tableData.title} />
</div>