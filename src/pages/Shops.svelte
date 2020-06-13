<script lang="typescript">
    import CommonTable from '../common/components/CommonTable/CommonTable.svelte'
    import SimpleElement from '../common/components/SimpleElement/SimpleElement.svelte'
	  import { onMount } from 'svelte';
    
    let data = [];

    const config = [
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
    ]

    onMount(async () => {
        setTimeout(() => {
          data = [
            { id: 1, name: 'name1', number: 1 },
            { id: 2, name: 'name2', number: 2 },
            { id: 3, name: 'name3', number: 3 },
            { id: 4, name: 'name4', number: 4 },
            { id: 5, name: 'name5', number: 5 },
          ]
        }, 10000)
    });
</script>

<style lang="scss">
</style>

<div>
    <CommonTable data={data} config={config} />
</div>