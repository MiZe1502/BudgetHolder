<script lang="typescript">
    import CommonTable from "../common/components/CommonTable/CommonTable.svelte"
    import SimpleElement from '../common/components/SimpleElement/SimpleElement.svelte'

    const data = [
      { id: 1, name: 'name1', number: 1 },
      { id: 2, name: 'name2', number: 2 },
      { id: 3, name: 'name3', number: 3 },
      { id: 4, name: 'name4', number: 4 },
      { id: 5, name: 'name5', number: 5 }
    ]

    const config = [
        {
            header: "id",
            component: SimpleElement,
        },
        {
            header: "name",
            component: SimpleElement,
        },
        {
            header: "number",
            component: SimpleElement,
        },
    ]
</script>

<style lang="scss">
</style>

<div>
    <CommonTable data={data} config={config} />
</div>