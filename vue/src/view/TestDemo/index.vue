<template>
  <!-- 控制dialog -->
  <el-button type="primary" size="default" @click="isShow = true">按钮</el-button>

  <div class="i-logos-vue"></div>

  <!-- 自己写的dialog组件 -->
  <DiaLog :isShow="isShow" :isScroll="true" @close="isShow = false">
    <template #title>这是一个标题</template>
    <template #default>这是一个内容</template>
    <template #bottom>
      <button>这是一个按钮</button>
    </template>
  </DiaLog>



  <!-- 测试pinia -->
  <div>{{ globalStore.count }}</div>
  <div>{{ globalStore.doubleCount() }}</div>
  <el-button type="primary" plain @click="globalStore.countAdd">+1</el-button>

  Lorem ipsum dolor sit amet consectetur adipisicing elit. Nisi odio incidunt provident est nobis natus nihil tempora
  repellat facere ducimus neque id, veritatis architecto non, maiores ab blanditiis. Perspiciatis omnis debitis
  tempore
  necessitatibus architecto, a quam repellat doloribus ducimus voluptate sed consectetur iure molestiae inventore
  commodi non! Eveniet explicabo officiis unde numquam inventore, quo quis possimus iste vel molestiae porro odit
  blanditiis dignissimos est voluptas nesciunt laudantium recusandae earum in tenetur doloribus assumenda ea qui. Quam
  voluptas cumque eos! Blanditiis, odio nobis temporibus optio laborum, similique cupiditate iusto vel doloribus id
  praesentium dolor fugiat officia ab qui rem accusamus quas? Voluptatem tempora officia adipisci voluptate, possimus
  totam expedita, commodi cum laborum explicabo, nesciunt accusamus odit quam dolorum. Tempora veniam aut molestias
  expedita impedit nam accusamus adipisci consequatur id ullam dignissimos, recusandae numquam officia ipsum
  aspernatur
  velit. Corrupti maiores quia accusantium, necessitatibus corporis, assumenda officia magni dolore unde odit quo
  aspernatur magnam explicabo, voluptatem natus consequatur commodi neque aut fuga repellendus quaerat rerum
  voluptates
  mollitia. Sapiente libero repellat nisi consequatur autem, praesentium quasi enim quia blanditiis unde doloribus
  iste
  sit explicabo nostrum consequuntur veritatis. Perferendis labore reprehenderit iure tempora illo quas dicta ipsum ea
  molestiae sit! Voluptas dolorum porro laboriosam earum asperiores dicta veniam necessitatibus et sapiente quidem
  odio
  consequatur ratione ab maxime, distinctio rem quibusdam in provident nam eaque doloribus. Aut soluta debitis odit
  aspernatur, corrupti consequuntur hic suscipit, eos cupiditate enim reprehenderit consequatur veritatis libero
  reiciendis. Maxime tempore et voluptatibus libero illum debitis ut nulla! Minus dolores vel, quam recusandae
  voluptatibus laudantium quasi, blanditiis aliquam exercitationem similique dolor quo fugiat deserunt velit adipisci.
  Officiis provident quaerat error cupiditate earum ex temporibus magni rem nihil alias vel aliquam, distinctio modi
  fugiat mollitia ab deleniti impedit. Placeat est, molestias asperiores provident, consequuntur itaque illum
  reiciendis, fugit facere assumenda veritatis quibusdam quisquam nam qui tempore vitae corporis! Nulla quam quibusdam
  ullam nihil ex laudantium, hic cum suscipit inventore similique itaque illo vitae ipsa ratione architecto eligendi
  eaque, dicta saepe quae, nesciunt consectetur deserunt reprehenderit reiciendis possimus. Nisi dolores, reiciendis
  velit ut laborum ullam quod eligendi itaque expedita libero illum asperiores vero recusandae necessitatibus maxime
  perferendis in incidunt corrupti magni molestias quidem! Itaque, officiis. Facilis nisi sint autem ipsam assumenda
  ab
  maxime molestias eum? Qui nesciunt temporibus maiores obcaecati, magni, sit hic est totam incidunt voluptas unde rem
  numquam reprehenderit nam sequi voluptatem mollitia cupiditate voluptatibus! Ipsa voluptatibus nesciunt, eum omnis
  voluptas, reiciendis quaerat laboriosam a tempora consequuntur optio sequi ullam aliquid soluta voluptatem deserunt
  accusantium. Commodi amet molestiae, labore in quo assumenda quasi dolorum sint, aut, quidem id exercitationem velit
  illo incidunt molestias magni accusamus nihil modi dolore tempora sequi repudiandae veritatis itaque vero. Adipisci
  quisquam et maiores obcaecati odit labore ad iste itaque facere, ipsam aliquam rerum, sequi, fugit vel error
  quibusdam
  iusto totam vitae porro? Aspernatur et mollitia unde ex. Voluptates facere quo, dolorum nulla corporis minus
  explicabo
  molestias autem, laborum deleniti optio accusamus laboriosam, odio nostrum. Hic, reprehenderit earum asperiores ab
  minus necessitatibus id nesciunt unde esse suscipit doloribus reiciendis in dolores mollitia ratione nemo pariatur
  quis at, modi nobis assumenda libero tempore totam excepturi! Molestias provident mollitia animi rem praesentium ab
  laudantium qui maxime quos neque est pariatur quas dolorem aperiam sunt eius recusandae cupiditate officiis, facilis
  voluptatem optio? Fugiat, laboriosam. Error officia doloremque culpa fuga ea exercitationem quod eum repudiandae
  rerum
  eveniet et quae earum debitis, minima, obcaecati maiores delectus, dolorem laboriosam suscipit reprehenderit
  repellendus maxime! Veniam necessitatibus voluptates repellendus! Facere possimus suscipit ipsum quisquam cupiditate
  consectetur numquam eum distinctio corporis totam sapiente ut neque molestias accusantium illum sit animi, ad ipsa
  fugiat repellendus perspiciatis deleniti? Praesentium doloribus quae unde expedita. Maxime adipisci enim ex? Earum
  rem
  quas sequi aliquid iste harum sed reiciendis? Alias quidem temporibus, sed, quisquam doloremque labore, est nemo
  veniam commodi deleniti exercitationem molestiae quas laudantium repudiandae iusto! Modi beatae laboriosam nisi est
  alias facilis veritatis dolore quas. Officia dolor nesciunt voluptatem unde aliquid inventore corrupti quisquam
  impedit consequuntur provident. Voluptates vitae voluptatibus pariatur ipsam corrupti possimus aliquid aperiam
  explicabo, repellat quisquam iure ratione obcaecati sequi architecto illo rerum voluptate impedit nisi blanditiis
  sunt, reiciendis consequuntur laboriosam. Blanditiis porro ab sint eius, amet deserunt earum temporibus tempora
  sequi
  non voluptates, soluta aspernatur debitis explicabo doloribus, nemo vero. Voluptatibus, eos impedit quaerat
  aspernatur
  maxime assumenda fugiat harum nam quisquam sint consequuntur ducimus totam officiis optio illo alias consectetur,
  sapiente vel minus quidem repellat eaque blanditiis. Eius harum voluptatibus libero corrupti modi commodi, nam illo
  at
  a iure non dolorem quod vel animi id? Exercitationem commodi culpa aperiam non sit possimus sint omnis quisquam ut
  iusto perferendis eveniet vitae quaerat officiis numquam nobis nostrum, optio voluptatibus ex accusantium tempore?
  Quas saepe aspernatur quam maiores voluptatibus deserunt ad facere deleniti eius. Ipsa accusantium, blanditiis ipsum
  ex quos, reprehenderit quod et cupiditate incidunt eligendi, nam cumque consequuntur eveniet fugiat nulla commodi
  aut!
  Magnam suscipit dolor hic ad praesentium eos eum, iure harum, fugit expedita ipsum nihil vero veritatis minima odio
  ex
  commodi porro ab, odit nulla. Aspernatur eaque sunt iste illum recusandae suscipit eligendi minima repudiandae
  cupiditate odit soluta delectus sapiente accusantium nisi, officia dolorum beatae facilis blanditiis? Delectus,
  quisquam a ex necessitatibus adipisci totam accusamus quos iure deserunt modi animi veniam. Iure similique ad
  debitis,
  rem animi officiis et, ullam magnam nisi adipisci iste tempora sequi doloremque sit hic autem, neque accusamus
  itaque.
  Officiis animi mollitia quaerat quae incidunt! Vero veniam accusamus dolorum sapiente aliquam modi delectus omnis
  nulla, nisi vitae mollitia dolore consequatur harum asperiores adipisci fugiat, ducimus voluptatem! Repellat
  delectus,
  iure quos maxime reprehenderit veritatis ea amet aliquid suscipit iste. Animi error voluptatum voluptas quae, vel
  soluta vitae repudiandae reprehenderit eius corrupti incidunt officiis, sit suscipit totam porro distinctio
  consequatur est culpa eligendi, hic aut alias nostrum? Alias debitis quae unde consequuntur. Tenetur aspernatur
  expedita sit dolore, praesentium voluptatum ipsam officiis eum corporis consectetur et a nemo exercitationem quaerat
  quos architecto beatae voluptates cupiditate possimus modi vero fugiat. Provident quae voluptatibus sed facilis,
  vitae
  ullam et a, facere dolorem ad ea harum veritatis rem!

</template>

<script lang='ts' setup>
const isShow = ref<boolean>(false)
const globalStore = GlobalStore()

</script> 

<style lang='less' scoped>
</style>